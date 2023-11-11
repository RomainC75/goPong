import { TextField, Button } from '@mui/material'
import { alpha, styled } from '@mui/material/styles'

const color1 = '#635FC7'

export const PurpleTextField = styled(TextField)(({ theme }) => ({
  '& .MuiOutlinedInput-root': {
    '&:hover fieldset': {
      borderColor: alpha(theme.palette.primary.main, 0.25),
    },
    '&.Mui-focused fieldset': {
      borderColor: color1,
    },
  },
  '& label.Mui-focused': {
    color: color1,
  },
}))

export const PurpleButton = styled(Button)(() => ({
  color: 'white',
  backgroundColor: color1,
  '&:hover': {
    backgroundColor: alpha(color1, 0.75),
  },
}))
