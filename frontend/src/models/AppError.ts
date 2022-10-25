export class AppError extends Error {
  code?: string
  extra?: any
  type?: string
  handled = false
}
