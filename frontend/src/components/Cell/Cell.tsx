import styled from 'styled-components'
import colors from 'theme/colors'

export default styled.div`
  background-color: white;
  width: 100%;
  min-height: 80px;
  border-radius: 15px;
  display: flex;
  box-shadow: 1px 1px 1px #48529944;
  overflow: hidden;
  &:hover {
    background-color: ${colors.adminBlue}20;
    cursor: pointer;
  }
`
