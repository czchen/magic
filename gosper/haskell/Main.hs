import Data.List.Split
import System.Environment
import Text.Read

import Gosper

genList :: Integer -> Integer -> [Integer] -> [Integer]
genList item count list
    | count <= 0 = list
    | otherwise = case getNext item of
        Nothing -> list
        Just next -> item: genList next (count - 1) list

main = do
    args <- getArgs
    if length args /= 2
        then
            putStrLn "program [first] [count]"
        else
            let
                (first:count:_) = map (\x -> readMaybe x :: Maybe Integer) args
                in
                    case (first, count) of
                        (Just first, Just count) -> putStrLn $ show $ genList first count []
                        (_, _) -> putStrLn "[first] and [count] must be integer"
