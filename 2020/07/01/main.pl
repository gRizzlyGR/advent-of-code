:- consult('kb.pl').

holds(A, B, _) :- contains(A, B, _).
holds(A, B, _) :- contains(A, X, _), holds(X, B, _).

% Part 1
count_holders(Contained, Count) :- 
    setof(Container, holds(Container, Contained, _), Containers),
    length(Containers, Count).