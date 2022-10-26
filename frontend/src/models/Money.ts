export enum Currency {
  EUR = 'EUR',
}

export type MoneyValue = {
  amount: number
  currency: Currency
}

function currencySymbol(cur: Currency): string {
  return { [Currency.EUR]: '€' }[cur]
}

export function stringifyMoney(value: MoneyValue): string {
  return `${value.amount / 100}${currencySymbol(value.currency)}`
}
