// 1 digit
// 1 special charactere
// 1 upper case

export const isPasswordValidFn = (pass: string): boolean => {
  return (
    isSpecialCharPresent(pass) &&
    isMoreThan8Characters(pass) &&
    isUpperCasePresent(pass) &&
    isNumericPresent(pass)
  )
}

const isNumericPresent = (pass: string): boolean => {
  const match: string[] | null = pass.match(/\d/g)
  return match && match.length ? true : false
}

const isUpperCasePresent = (pass: string): boolean => {
  const lowerPass = pass.toLowerCase()
  return lowerPass === pass ? false : true
}

const isMoreThan8Characters = (pass: string): boolean => {
  return pass.length > 7 ? true : false
}

const isSpecialCharPresent = (pass: string): boolean => {
  const match: string[] | null = pass.match(/[*+,-./:;()<=>?@]/g)
  return match ? true : false
}

export const isEmailValidFn = (pass: string): boolean => {
  const match: string[] | null = pass.match(/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g)
  return match ? true : false
}

export const detailsAboutNeededCharactersInPass = (pass: string): string => {
  let str = 'need at least '
  if (!isMoreThan8Characters(pass)) {
    str += '8 characters including : '
  }
  if (!isNumericPresent(pass)) {
    str += '1digit, '
  }
  if (!isUpperCasePresent(pass)) {
    str += '1 upper case, '
  }
  if (!isSpecialCharPresent(pass)) {
    str += '1 special char []*+,-./:;()<=>?@'
  }
  str += ')'
  return str
}
