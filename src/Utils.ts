
export default class Utils {
  static formatMonth(dt: Date) {
    return `${dt.getFullYear()}/${('00' + (dt.getMonth() + 1)).slice(-2)}`
  }

  static formatDate(dt: Date) {
    return `${dt.getFullYear()}/${('00' + (dt.getMonth() + 1)).slice(-2)}/${('00' + dt.getDate()).slice(-2)}`
  }
}
