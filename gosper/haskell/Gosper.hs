module Gosper
( next
) where

import Data.Bits

next :: Integer -> Integer
next x =
    let
        y = x .&. (- x)
        c = x + y
    in
        (((x `xor` c) `shiftR` 2) `div` y) .|. c
