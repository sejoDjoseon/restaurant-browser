import Container from 'components/Container/Container'
import styled from 'styled-components'
import colors from 'theme/colors'

export default styled(Container)`
  width: 100%;
  min-height: 80px;
  display: flex;
  &:hover {
    background-color: ${colors.adminBlue}20;
    cursor: pointer;
  }
`
