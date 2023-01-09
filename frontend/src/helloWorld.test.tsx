import helloWorld from './helloWorld';

test('helloWorld returns "Hello World!"', () => {
    expect(helloWorld()).toBe('Hello World!');
});