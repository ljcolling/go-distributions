## About

This repo contains functions for (the density of) the following statistical distributions:

- Uniform (`Dunif`)
- Binomial (`Bbinom`) 
- Beta (`Dbeta`) 
- scaled and shifted t-distribution (`Scaled_shifted_t`)
- Normal distribution (`Dnorm`) 
- Non-central t distribution (`Dt`)
- Cauchy distribution (`Dcauchy`)

And a helper function:

- Integrate

Apart from the `Dt`, all the distributions are simply wrappers around the functions provided by [GoNum](https://www.gonum.org) (The cauchy distribution is implemented as a 1 df student t distribution). For the non-central t distribution, a standard algorithm that is implemented in, for example, SciPy, is used. 


The functions written so that their interface is consistent with that employed by the corresponding functions in `R`.


