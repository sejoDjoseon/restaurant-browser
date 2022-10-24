import styled from "styled-components";
import COLORS from "theme/colors";

const Header = styled.div`
  background-color: ${COLORS.adminBlue};
  min-height: 10vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 2vmin);
  color: white;
`;

export default Header;