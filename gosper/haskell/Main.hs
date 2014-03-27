import Data.List.Split
import System.Environment

import Gosper

genList :: Integer -> Integer -> [Integer] -> [Integer]
genList item count list
    | count == 0 = list
    | otherwise  = item: genList (next item) (count - 1) list

main = do
    args <- getArgs
    let
        (item:count:_) = map (\x -> read x :: Integer) args
    putStrLn $ show $ genList item count []
