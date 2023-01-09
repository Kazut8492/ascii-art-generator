import helloWorld from '../random/helloWorld';

test('helloWorld returns "Hello World!"', () => {
    expect(helloWorld()).toBe('Hello World!');
});