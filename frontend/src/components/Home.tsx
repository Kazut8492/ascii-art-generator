import React, {useState} from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

const Home: React.FC = () => {
    const [input, setInput] = useState('');
    const [output, setOutput] = useState('');

    const handleFormChange = (event: React.FormEvent) => {
        // set value to the value of the input
        setInput((event.target as HTMLInputElement).value)
    }

    const handleFormSubmit = (event: React.FormEvent) => {
        event.preventDefault()
        // console.log('button clicked!')
        // console.log(value)

        // send the value to the backend
        fetch('http://localhost:8080/ascii', {
            method: 'POST',
            // mode: "cors",
            // cache: "no-cache",
            // credentials: "include",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(input)
        })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            setInput('')
            setOutput(data)
        })
        .catch(error=>console.log(error))
    }

    // Component to convert line breaks to <br /> tags
    const MultiLineBody = ({ body }: { body: string }) => {
        const texts = body.split('\n').map((item, index) => {
          return (
            <React.Fragment key={index}>
              {item}
              <br />
            </React.Fragment>
          );
        });
        return <div>{texts}</div>;
      };

    return (<>
        <Form onSubmit={handleFormSubmit}>
            <Form.Group className='mb-3' controlId='asciiInput'>
                <Form.Label>Type in what you want to convert into ascii art</Form.Label>
                <Form.Control type='text' placeholder='type in here' value={input} onChange={handleFormChange}/>
            </Form.Group>
            <Button type='submit'>Submit</Button>
        </Form>
        {/* {output && <MultiLineBody body={output} />} */}
        {output && <p className='ascii-output'>{output}</p>}
    </>)
}

export default Home