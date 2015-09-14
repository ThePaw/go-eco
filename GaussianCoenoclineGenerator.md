# Gaussian Coenocline Generator #

Some examples of behavior of the ccline\_gauss command.


# default settings, changing random number seed #

ccline\_gauss

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -z 2

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-02.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-02.png)

ccline\_gauss -z 3

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-03.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-03.png)

ccline\_gauss -z 4

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-04.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-04.png)

ccline\_gauss -z 5

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-05.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-05.png)


# changing variance of population size #

ccline\_gauss

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -ps 0.5

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-06.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-06.png)

ccline\_gauss -ps 0.8

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-07.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-07.png)

ccline\_gauss -ps 1.5

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-08.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-08.png)

# changing mean tolerance #

ccline\_gauss

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -tm 0.18

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-10.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-10.png)

ccline\_gauss -tm 0.3

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-11.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-11.png)


# changing variance of tolerance #

ccline\_gauss

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -ts 0.5

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-13.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-13.png)

ccline\_gauss -ts 0.55

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-14.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-14.png)

ccline\_gauss -ts 1.2

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-15.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-15.png)

# adding noise #
ccline\_gauss

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -e 0.1

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-16.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-16.png)

ccline\_gauss -e 0.2

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-17.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-17.png)

ccline\_gauss -e 0.5

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-18.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-18.png)


# effect of sampling model #

ccline\_gauss	(regular spacing of samples)

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-01.png)

ccline\_gauss -s 1 	(uniform random spacing of samples)

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-19.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-19.png)

ccline\_gauss -s 2 	(Poisson point process spacing of samples)

![http://www.gli.cas.cz/home/cejchan/go/cc/cc-20.png](http://www.gli.cas.cz/home/cejchan/go/cc/cc-20.png)

