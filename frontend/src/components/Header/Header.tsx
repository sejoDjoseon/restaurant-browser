import styled from 'styled-components'
import COLORS from 'theme/colors'

const Header = styled.div`
  background-color: ${COLORS.adminBlue};
  height: 8vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 1.5vw);
  color: white;
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1;
`

export default Header
