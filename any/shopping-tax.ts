function shoppingTax(pricing: number, taxRate: number = 0.25): number {
  const taxAmount = pricing * taxRate;
  return pricing + taxAmount;
}

console.log(shoppingTax(10000)); // 125
console.log(shoppingTax(20000, 0.1)); // 220 (with 10% tax)
