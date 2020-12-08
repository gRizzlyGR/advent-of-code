# Knowledge Base Structure
Starting from this fact:
```
bright gray bags contain 2 bright gold bags, 5 dull lavender bags.
```

we generate atoms like this:
```prolog
contains(bag(color(bright_gray)), bag(color(bright_gold)), 2).

contains(bag(color(bright_gray)), bag(color(dull_lavender)), 5).
```

i.e the bag container, the bag contained, and how many bags of the same kind are contained.

# How-To

First, create the knowledge base:
```bash
go run main.go < input > kb
```

Then, fire up prolog:
```bash
swipl -f main.pl
```

For the first part use this command:
```prolog
count_holders(bag(color(shiny_gold)), C).
```

For the second part, use this command:
```prolog
count_held_with_complessive_qty(bag(color(shiny_gold)), C).
```

# Notes
I've developed the program with these specs:
- Ubuntu 16.04
- Golang 1.13 (for the knowledge base generation)
- SWI-Prolog v. 7.2.3