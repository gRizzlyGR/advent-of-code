:- consult('kb.pl').

holds(A, B) :- contains(A, B, _).
holds(A, B) :- contains(A, X, _), holds(X, B).

% Part 1: How many holders, direct and indirect
count_holders(Contained, Count) :- 
    setof(Container, holds(Container, Contained), Containers),
    length(Containers, Count).

holds_with_qty(A, B, Qty) :- contains(A, B, Qty).
holds_with_qty(A, B, Qty) :- 
    contains(A, X, PartialQty1),
    holds_with_qty(X, B, PartialQty2),
    Qty is PartialQty1 * PartialQty2.

% Part 2: What is the total of all quantities for each item held (directly and indirectly)
count_held_with_complessive_qty(Container, Total) :-
    findall(Quantity, holds_with_qty(Container, _, Quantity), Quantities),
    sum(Quantities, Total).

sum([Sum], Sum).
sum([Head | Tail], Sum) :-
    sum(Tail, PartialSum),
    Sum is PartialSum + Head.
