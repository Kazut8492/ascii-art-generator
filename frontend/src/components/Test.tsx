import React from 'react'

interface Props {
    name: string;
    age: number;
}

const Test: React.FC<Props> = (props) => {
    return (<div>
        Hello, this is {props.name} and I am {props.age} years old.
    </div>)
}

export default Test;