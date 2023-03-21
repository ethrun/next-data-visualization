import React, { useState } from "react";
import styled from "styled-components";

const StyledInputWrapper = styled.div`
  z-index: 1;
`;

const StyledTitle = styled.span`
  display: flex;
  justify-content: center;
  align-items: center;
  font-weight: 500;
  gap: 10px;
  font-size: 16px;
`;

const Button = styled.button`
  color:#ffffff
`


export const DatabaseInput = () => {
    const [value, setValue] = useState('')
    const saveUrl = () => {
        localStorage.setItem('url', value)
    }
    return (
        <StyledInputWrapper>
            <StyledTitle>
                Database url:
                <input value={value} onChange={(e) => { setValue(e.target.value) }}>
                </input>
                <Button onClick={saveUrl}>Save</Button>
            </StyledTitle>
        </StyledInputWrapper>
    );
};
