# exdraw
exdraw is an external drawing package based on [tfriedel6/canvas](https://github.com/tfriedel6/canvas) and [go-gl/glfw](https://github.com/go-gl/glfw ). 

It allows you to create windows in **goroutines** and draw in them, similar to the **HTML5** canvas API.

# Why
At one time, we found out that other packages block the main **goroutine** and do not provide the functionality we need, which is contrary to our tasks.
Therefore, we solved this problem in our own way.

# Community
You can contribute to this package.

# Examples
Look in the examples folder for some drawing examples.