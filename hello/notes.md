## Writting tests

Writting a test is just like writting a function, with a few rules:

- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail. 

**t.Errorf**

We are calling the **Errorf** method on our t which will print out a message and fail the test. The f stands for format which allows us to build a string with values inserted into the placeholder values %q. When you made the test fail it should be clear how it works.




## Discipline
Let's go over the cycle again
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

On the face of it this may seem tedious but sticking to the feedback loop is important.
Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with the safety of tests.
Seeing the test fail is an important check because it also lets you see what the error message looks like. As a developer it can be very hard to work with a codebase when failing tests do not give a clear idea as to what the problem is.
By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.
By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run.



## The TDD process and why steps are so important
- Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
- Writing the smallest amount of code to make it pass so we know we have working software
- Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

In our case we've gone from **Hello()** to **Hello("name")**, to Hello**("name", "French")** in small, easy to understand steps.


This is of course trivial compared to "real world" software but the principles still stand. TDD is a skill that needs practice to develop, but by breaking problems down into smaller components that you can test, you will have a much easier time writing software.