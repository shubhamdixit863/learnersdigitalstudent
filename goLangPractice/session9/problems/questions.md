You are tasked with designing a plugin-based payment gateway system that allows multiple payment providers
(e.g., PayPal, Stripe, Razorpay) to integrate dynamically. The system should:

1. Support multiple payment providers with different implementations.

2. Enforce a standard interface so that all providers follow a common structure.

3. Allow dynamic registration of providers without modifying core logic.

4. Use type assertions to extract provider-specific details when needed.

5. Ensure proper method implementations for processing payments and refunds.

Instructions



1. Define a Payment Interface

• Create a PaymentProcessor interface that requires the following methods:

• Pay(amount float64) string → Processes a payment.

• Refund(transactionID string) string → Refunds a transaction.

• GetProviderName() string → Returns the provider’s name.



2. Implement Payment Providers

• Create three struct implementations: PayPal, Stripe, and Razorpay.

• Each struct should:

• Store unique details (e.g., ClientID, APIKey, MerchantID).

• Implement the PaymentProcessor interface.



3. Create a Payment Gateway Manager

• Implement a PaymentGateway struct that:

• Stores registered payment providers dynamically.

• Has methods to:

• RegisterProvider(provider PaymentProcessor) → Adds a new payment provider.

• ProcessPayment(providerName string, amount float64) (string, error) → Finds the provider and processes a payment.

• IssueRefund(providerName, transactionID string) (string, error) → Processes a refund.



4. Use Type Assertions

• Implement a function that:

• Takes a PaymentProcessor and extracts provider-specific details using type assertions.

• Prints unique attributes (e.g., ClientID, APIKey, MerchantID).



5. Test the System

• Register multiple providers (PayPal, Stripe, Razorpay).

• Process a payment using one of the providers.

• Issue a refund.

• Use type assertion to extract and display provider-specific details.


6. Create a custome error for different payment methods

• Add return errors from methods, if the operation fails as the second argument

• internal/services

• internal/utils/errors.go//Paypal error

