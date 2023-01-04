import React, {useState} from 'react'
import TextField from "@mui/material/TextField";

const searchBar = () => {
    const [inputText, setInputText] = useState("");
    let inputHandler = (e: { target: { value: string; }; }) => {
        //convert input text to lower case
        var lowerCase = e.target.value.toLowerCase();
        setInputText(lowerCase);
    };

}


export default searchBar;