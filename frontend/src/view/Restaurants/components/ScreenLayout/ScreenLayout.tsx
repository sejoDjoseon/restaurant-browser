import styled from 'styled-components'

export const ScreenContainer = styled.div`
  display: flex;
`
export const SectionContainer = styled.div<{ widthVW: number }>`
  width: ${({ widthVW }) => widthVW}vw;
`
