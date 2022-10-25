import styled from 'styled-components'
import COLORS from 'theme/colors'

const Header = styled.div`
  background-color: ${COLORS.adminBlue};
  height: 8vh;
  min-height: 80px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 1.5vw);
  color: white;
`

export default Header
