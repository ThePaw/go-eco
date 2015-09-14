# Plotting a CSV matrix rowwise #

plt plots data matrix ("data.csv") rows as lines.

This is a toy project based on [plotinum](https://code.google.com/p/plotinum/) pkg and, when mature, probably belongs elsewhere. For the time being, I use it here to visualize generated coenoclines.
It uses selected colors (dark ones) from the [W3C palette](http://www.w3.org/TR/SVG/types.html#ColorKeywords). Complete palette also included with the [plt pkg](https://code.google.com/p/go-eco/source/browse/#hg%2Feco%2Fplt).


## Simple usage ##

Just put your data into "data.csv" and type:

> plt  < data.csv

and inspect the "plot.svg" file.

![http://www.gli.cas.cz/home/cejchan/go/plot1.png](http://www.gli.cas.cz/home/cejchan/go/plot1.png)

## Flags ##

  * **-h** help
  * **-o** _plot.svg_: output file (type recognized by extension: SVG [default](default.md) , PDF, PNG)
  * **-t** _: plot title
  * **-x**_: label of the X axis
  * **-y** _: label of the Y axis
  * **-p**_: name of the palette to be used

## Examples ##

plt -t Plot -x X -y Y -p Gold -o plot2.png  < data.csv

![http://www.gli.cas.cz/home/cejchan/go/plot2.png](http://www.gli.cas.cz/home/cejchan/go/plot2.png)

plt -t 'Example coenocline' -x Altitude -y 'Population density' -p Hilite -o plot3.png  < data.csv

![http://www.gli.cas.cz/home/cejchan/go/plot3.png](http://www.gli.cas.cz/home/cejchan/go/plot3.png)