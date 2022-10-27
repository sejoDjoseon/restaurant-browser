import styled from 'styled-components'

export const ScreenContainer = styled.div`
  display: flex;
`
export const SectionContainer = styled.div<{ widthVW: number }>`
  width: ${({ widthVW }) => widthVW}vw;
`
export const ItemList = styled.div`
  display flex;
  flex-direction: column;
  padding: 30px;
  align-items: center;
`
