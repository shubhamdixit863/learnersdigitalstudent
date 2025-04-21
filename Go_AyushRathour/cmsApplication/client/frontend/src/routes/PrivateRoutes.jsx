import { Navigate, useLocation } from "react-router-dom";

const PrivateRoute = ({ children, isAdmin = false }) => {
  const token = localStorage.getItem(isAdmin ? "adminToken" : "token");
  const location = useLocation();

  return token ? children : <Navigate to={isAdmin ? "/admin/login" : "/"} state={{ from: location }} />;
};

export default PrivateRoute;
