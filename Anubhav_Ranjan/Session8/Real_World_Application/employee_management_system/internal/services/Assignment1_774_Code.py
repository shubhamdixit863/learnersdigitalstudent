import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
import random
from collections import Counter
import heapq

# Set random seed for reproducibility
random.seed(42)
np.random.seed(42)

# Function to create a random connected graph
def generate_random_connected_graph(n_vertices, edge_probability=0.0003):
    # Start with a minimum spanning tree to ensure connectivity
    G = nx.path_graph(n_vertices)

    # Add random edges with the given probability
    for i in range(n_vertices):
        for j in range(i+2, n_vertices):  # Skip adjacent nodes (already connected)
            if random.random() < edge_probability:
                G.add_edge(i, j)

    # Ensure the graph is connected
    if not nx.is_connected(G):
        print("Warning: Generated graph was not connected. Connecting components...")
        components = list(nx.connected_components(G))
        for i in range(len(components) - 1):
            # Connect each component to the next one
            u = random.choice(list(components[i]))
            v = random.choice(list(components[i + 1]))
            G.add_edge(u, v)

    print(f"Generated graph with {n_vertices} vertices and {G.number_of_edges()} edges")
    return G

# Function to calculate normalized degree centrality
def normalized_degree_centrality(G):
    n = G.number_of_nodes()
    max_degree = n - 1  # Maximum possible degree in a simple graph

    degree_centrality = {}
    for node in G.nodes():
        degree_centrality[node] = G.degree(node) / max_degree

    return degree_centrality

# Function to calculate closeness centrality
def closeness_centrality(G):
    n = G.number_of_nodes()
    closeness = {}

    # Calculate shortest paths between all pairs of nodes
    shortest_paths = dict(nx.all_pairs_shortest_path_length(G))

    for node in G.nodes():
        total_distance = sum(shortest_paths[node].values())
        # Closeness centrality is the inverse of the average shortest path length
        closeness[node] = (n - 1) / total_distance if total_distance > 0 else 0

    return closeness

# Function to calculate betweenness centrality
def betweenness_centrality(G):
    n = G.number_of_nodes()
    betweenness = {node: 0 for node in G.nodes()}

    # Calculate shortest paths between all pairs of nodes
    for s in G.nodes():
        # Use BFS to find all shortest paths from s to all other nodes
        visited = set()
        queue = [(s, [s])]

        # Dictionary to store shortest paths from s to each node
        shortest_paths = {s: [[s]]}

        while queue:
            node, path = queue.pop(0)
            if node not in visited:
                visited.add(node)

                for neighbor in G.neighbors(node):
                    if neighbor not in visited:
                        new_path = path + [neighbor]
                        if neighbor not in shortest_paths:
                            shortest_paths[neighbor] = [new_path]
                            queue.append((neighbor, new_path))
                        elif len(new_path) == len(shortest_paths[neighbor][0]):
                            shortest_paths[neighbor].append(new_path)
                            queue.append((neighbor, new_path))

        # Count shortest paths passing through each node
        for t in G.nodes():
            if t != s:
                for path in shortest_paths.get(t, []):
                    for node in path[1:-1]:  # Exclude source and target
                        betweenness[node] += 1

    # Normalize betweenness centrality
    total_paths = sum(betweenness.values())
    if total_paths > 0:
        for node in betweenness:
            betweenness[node] /= total_paths

    return betweenness

# Novel idea 1: Community influence score
def community_influence_score(G):
    # Detect communities using the Louvain method
    communities = nx.community.louvain_communities(G)

    # Map each node to its community
    node_to_community = {}
    for i, community in enumerate(communities):
        for node in community:
            node_to_community[node] = i

    # Calculate community influence scores
    influence_scores = {}
    for node in G.nodes():
        neighbor_communities = set()
        for neighbor in G.neighbors(node):
            neighbor_communities.add(node_to_community.get(neighbor, 0))

        # Score is the number of different communities the node connects to
        influence_scores[node] = len(neighbor_communities) / len(communities) if communities else 0

    return influence_scores

# Novel idea 2: Second-level betweenness refinement
def second_level_betweenness(G, top_nodes, original_betweenness):
    refined_betweenness = {node: 0 for node in G.nodes()}

    # Find all shortest paths involving top nodes
    for s in top_nodes:
        for t in G.nodes():
            if s != t:
                for path in nx.all_shortest_paths(G, s, t):
                    for node in path[1:-1]:  # Exclude source and target
                        refined_betweenness[node] += 1

    # Normalize the refined scores
    total = sum(refined_betweenness.values())
    if total > 0:
        for node in refined_betweenness:
            refined_betweenness[node] /= total

    # Combine with original betweenness centrality
    for node in G.nodes():
        refined_betweenness[node] = 0.7 * original_betweenness[node] + 0.3 * refined_betweenness[node]

    return refined_betweenness

# Novel idea 3: Information flow potential
def information_flow_potential(G, degree_centrality, closeness_centrality):
    flow_potential = {}

    for node in G.nodes():
        # Calculate the average centrality of neighbors
        neighbor_centrality = 0
        neighbors = list(G.neighbors(node))
        if neighbors:
            neighbor_centrality = sum(degree_centrality[neighbor] for neighbor in neighbors) / len(neighbors)

        # Flow potential combines node's own centrality with its neighbors' centrality
        flow_potential[node] = (0.6 * closeness_centrality[node] + 0.4 * neighbor_centrality)

    # Normalize the scores
    max_flow = max(flow_potential.values()) if flow_potential else 1
    for node in flow_potential:
        flow_potential[node] /= max_flow

    return flow_potential

# Combined centrality score
def combined_centrality(G, alpha=0.3, beta=0.3, gamma=0.4):
    # Calculate basic centrality measures
    degree_cent = normalized_degree_centrality(G)
    closeness_cent = closeness_centrality(G)
    betweenness_cent = betweenness_centrality(G)

    # Novel measures
    community_influence = community_influence_score(G)

    # Identify top nodes based on first-level computation
    first_level_scores = {}
    for node in G.nodes():
        first_level_scores[node] = (
            alpha * degree_cent[node] +
            beta * closeness_cent[node] +
            gamma * betweenness_cent[node]
        )

    # Get top 10% of nodes
    top_nodes_count = max(int(G.number_of_nodes() * 0.1), 10)
    top_nodes = set(heapq.nlargest(top_nodes_count, first_level_scores, key=first_level_scores.get))

    # Refine betweenness centrality
    refined_betweenness = second_level_betweenness(G, top_nodes, betweenness_cent)

    # Calculate information flow potential
    flow_potential = information_flow_potential(G, degree_cent, closeness_cent)

    # Combine all metrics into a final score
    combined_scores = {}
    for node in G.nodes():
        combined_scores[node] = (
            0.20 * degree_cent[node] +
            0.20 * closeness_cent[node] +
            0.25 * refined_betweenness[node] +
            0.20 * community_influence[node] +
            0.15 * flow_potential[node]
        )

    return combined_scores

# Function to visualize the graph with colored nodes based on centrality
def visualize_graph(G, centrality_scores, top_k=10, output_file="graph_visualization.png"):
    plt.figure(figsize=(12, 12))

    # Get top k nodes
    top_nodes = heapq.nlargest(top_k, centrality_scores, key=centrality_scores.get)

    # Create a smaller graph with top nodes and their neighbors for better visualization
    subgraph_nodes = set(top_nodes)
    for node in top_nodes:
        subgraph_nodes.update(G.neighbors(node))

    if len(subgraph_nodes) > 200:
        subgraph_nodes = set(random.sample(list(subgraph_nodes), 200))
        subgraph_nodes.update(top_nodes)

    H = G.subgraph(subgraph_nodes)

   
    pos = nx.spring_layout(H, seed=42)

    nx.draw_networkx_nodes(H, pos, node_color='lightgray', node_size=100, alpha=0.8)


    nx.draw_networkx_edges(H, pos, alpha=0.3)

    # Draw top k nodes with colors based on centrality
    colors = plt.cm.rainbow(np.linspace(0, 1, top_k))
    for i, node in enumerate(top_nodes):
        if node in H:
            nx.draw_networkx_nodes(H, pos, nodelist=[node], node_color=[colors[i]],
                                   node_size=300, alpha=1.0)

    labels = {node: f"{node}\n({centrality_scores[node]:.3f})" for node in top_nodes if node in H}
    nx.draw_networkx_labels(H, pos, labels=labels, font_size=8, font_weight='bold')

    sm = plt.cm.ScalarMappable(cmap=plt.cm.rainbow, norm=plt.Normalize(vmin=0, vmax=1))
    sm._A = []  # Fake array for the colorbar
    cbar = plt.colorbar(sm, ax=plt.gca())
    cbar.set_label('Centrality Score')

    plt.title("Graph Visualization with Top 10 Central Nodes")
    plt.axis('off')
    plt.tight_layout()
    plt.savefig(output_file, dpi=300, bbox_inches='tight')
    plt.close()

    plt.figure(figsize=(12, 12))

    pos = nx.spring_layout(G, seed=42, k=0.3, iterations=50)

    node_colors = ['lightgray'] * G.number_of_nodes()
    node_sizes = [10] * G.number_of_nodes()

    for i, node in enumerate(top_nodes):
        node_colors[node] = colors[i]
        node_sizes[node] = 100

    nx.draw(G, pos, node_color=node_colors, node_size=node_sizes,
            width=0.2, alpha=0.7, with_labels=False)

    plt.title("Full Graph with Top 10 Central Nodes Highlighted")
    plt.axis('off')
    plt.tight_layout()
    plt.savefig("full_graph_visualization.png", dpi=300, bbox_inches='tight')
    plt.close()

def generate_report(G, centrality_scores, top_k=10):
    report = "# Graph Centrality Analysis Report\n\n"

    report += "## Graph Statistics\n"
    report += f"- Number of vertices: {G.number_of_nodes()}\n"
    report += f"- Number of edges: {G.number_of_edges()}\n"
    report += f"- Average degree: {2 * G.number_of_edges() / G.number_of_nodes():.2f}\n"
    
    return report

def main():
    # Generate a random connected graph with 1000 vertices
    G = generate_random_connected_graph(1000, edge_probability=0.003)

    # Calculate combined centrality scores
    print("Calculating centrality scores...")
    centrality_scores = combined_centrality(G)

    # Get the top 10 nodes with highest centrality scores
    top_nodes = heapq.nlargest(10, centrality_scores, key=centrality_scores.get)
    print("\nTop 10 nodes with highest centrality scores:")
    for i, node in enumerate(top_nodes):
        print(f"{i+1}. Node {node}: {centrality_scores[node]:.4f}")

    # Visualize the graph
    print("\nGenerating graph visualizations...")
    visualize_graph(G, centrality_scores, top_k=10)

    # Generate a detailed report
    print("Generating detailed report...")
    report = generate_report(G, centrality_scores)

    # Save the report to a file
    with open("centrality_analysis_report.md", "w") as f:
        f.write(report)

    print("\nAnalysis complete. Results saved to:")
    print("- graph_visualization.png")
    print("- full_graph_visualization.png")
    print("- centrality_analysis_report.md")

if __name__ == "__main__":
    main()