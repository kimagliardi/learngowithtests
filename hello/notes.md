## Writting tests

Writting a test is just like writting a function, with a few rules:

- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail. 

**t.Errorf**

We are calling the **Errorf** method on our t which will print out a message and fail the test. The f stands for format which allows us to build a string with values inserted into the placeholder values %q. When you made the test fail it should be clear how it works.