module Gosper
( getNext
) where

import Data.Bits

getNext :: Integer -> Maybe Integer
getNext x
    | x <= 0 = Nothing
    | otherwise =
        let
            y = x .&. (- x)
            c = x + y
        in
            Just $ (((x `xor` c) `shiftR` 2) `div` y) .|. c
