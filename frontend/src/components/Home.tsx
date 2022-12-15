import React from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

const Home: React.FC = () => {

    const handleFormSubmit = (event: React.FormEvent) => {
        event.preventDefault()
        console.log('button clicked!')
    }

    return (<>
        <Form onSubmit={handleFormSubmit}>
            <Form.Group className='mb-3' controlId='asciiInput'>
                <Form.Label>Type in what you want to convert into ascii art</Form.Label>
                <Form.Control type='text' placeholder='type in here' />
            </Form.Group>
            <Button type='submit'>Submit</Button>
        </Form>
    </>)
}

export default Home