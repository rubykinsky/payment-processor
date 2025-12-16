export enum PaymentGateway {
  Stripe = 'stripe',
  PayPal = 'paypal',
  Square = 'square',
  AuthorizeNet = 'authorize.net',
}

export enum Currency {
  USD = 'USD',
  EUR = 'EUR',
  JPY = 'JPY',
  CNY = 'CNY',
}

export type Address = {
  street: string;
  city: string;
  state: string;
  country: string;
  postalCode: string;
}

export type BillingAddress = Address & {
  name: string;
}

export type ShippingAddress = Address & {
  companyName: string;
}

export type CreditCard = {
  number: string;
  expMonth: number;
  expYear: number;
  cvc: string;
}

export type PaymentMethod = {
  type: PaymentGateway;
  card?: CreditCard;
  billingAddress?: BillingAddress;
  shippingAddress?: ShippingAddress;
}

export type Payment = {
  amount: number;
  currency: Currency;
  paymentMethod: PaymentMethod;
}

export type OrderItem = {
  id: string;
  name: string;
  quantity: number;
  price: number;
}

export type Order = {
  id: string;
  items: OrderItem[];
  subtotal: number;
  total: number;
  payment: Payment;
}