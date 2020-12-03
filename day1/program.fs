// Day1 Solution in F# //

let array1 =
    [|
        1728
        1954
        1850
        1825
        1732
        1536
        1759
        1877
        1400
        1579
        1708
        1047
        1810
        558
        1132
        1608
        1857
        1756
        1834
        1743
        1888
        1660
        1642
        1726
        541
        1519
        1407
        1875
        1618
        1331
        1878
        1626
        1200
        1346
        1830
        1403
        1557
        1890
        1543
        823
        1435
        1903
        1377
        1931
        1885
        1422
        1411
        1563
        1818
        1643
        2004
        1364
        1446
        1071
        1699
        1140
        1617
        1974
        1758
        1537
        1980
        1709
        1812
        1178
        1822
        1648
        1517
        1477
        1935
        1848
        1534
        1734
        1484
        1985
        1485
        1963
        1329
        1809
        1380
        1552
        1895
        215
        1844
        1138
        1194
        1938
        1774
        1823
        684
        1948
        1941
        1062
        1550
        1602
        1920
        1391
        1666
        1327
        1791
        1721
        1928
        1805
        1574
        1658
        1467
        1852
        1924
        1679
        2008
        1989
        1719
        1884
        1776
        1806
        1750
        1897
        1781
        1667
        1544
        1100
        1838
        1839
        1744
        1715
        1481
        1480
        1548
        1707
        1362
        1681
        1616
        1956
        1639
        1911
        1655
        1685
        1670
        1789
        1571
        1661
        1647
        1379
        1522
        1965
        1482
        1158
        1970
        1945
        1384
        1535
        1383
        1613
        1511
        1896
        1784
        1513
        841
        1619
        1645
        1125
        1932
        1873
        639
        1657
        1554
        1979
        1516
        1995
        1899
        1347
        1175
        1918
        1872
        1559
        1094
        1423
        1883
        1846
        1394
        1488
        1343
        1905
        1914
        1578
        1943
        1388
        1286
        966
        1342
        1528
        1702
        1452
        1936
        2005
        1188
        1683
        1133
        447
        1072
        1893
     |]

let is2020 x y = x + y = 2020
let is2020Part2 x y z = x + y + z = 2020
   
let rec findProductOfTwoNumbersThatAddUpTo2020 (numberArray : int[]) position : int =    
    let item = numberArray.[position]
    let nextPosition = position + 1
    let result = Array.tryFind(fun elem -> (is2020 item elem)) numberArray
    match result with        
    | Some x -> printfn "%d %d" item result.Value; item * result.Value
    | None -> findProductOfTwoNumbersThatAddUpTo2020 numberArray nextPosition

let rec findProductOf3NumbersThatAddUpTo2020 (numberArray : int[]) positionA positionB : int =    
    if positionA >= numberArray.Length  //no matches on second position so change it and try again.
    then findProductOf3NumbersThatAddUpTo2020 numberArray 0 (positionB+1) 
    else //
    let itemA = numberArray.[positionA]
    let itemB = numberArray.[positionB]
    let nextPositionA = positionA + 1
    let result = Array.tryFind(fun elem -> (is2020Part2 itemA itemB elem)) numberArray
    match result with        
    | Some x -> printfn "%d %d %d" itemA itemB result.Value; itemA * itemB * result.Value
    | None -> findProductOf3NumbersThatAddUpTo2020 numberArray nextPositionA positionB //no matches so recurse on the first position
    
[<EntryPoint>]
let main argv =                  
    let answer = findProductOfTwoNumbersThatAddUpTo2020 array1 0
    printfn "Part1 Answer is %d" answer

    let answer2 = findProductOf3NumbersThatAddUpTo2020 array1 0 1
    printfn "Part2 Answer is %d" answer2
    
    0 // return an integer exit code
