# solartax
Back of the envelope estimate of income tax due on solar electricity generation

Default result: 7,005 Euro tax due

Usage:

```
$ ./solartax -h
Usage of ./solartax:
  -annualtaxallowance float
        Income tax allowance. Eg 200.0 for first 200 euro income every year to be untaxed (default 200)
  -incometaxrate float
        Income tax rate. For 52% use 0.52 here. (default 0.52)
  -selltogridfraction float
        Fraction of generated energy sold to grid. Eg if you export 2/3 of generated electricity and use 1/3, set to 0.66. If you use 2/3 and export 1/3, set to 0.33 (default 0.66)
  -solarpanellifetime float
        Lifespan of solar panels (duration of investment). Eg 25.0 for 25 years (default 25)
  -totalannualgeneration float
        Total amount of electricity generated in a year in kilowatts. Eg for 3.2 megawatts set to 3200 (default 3200)
  -unitbuyprice float
        Price for buying a unit from the grid. eg 0.31 for 31 cents per unit (default 0.31)
  -unitbuypriceinflation float
        Rate at which price for selling to the grid increases. Eg for 5% per year use 0.05 here (default 0.04)
  -unitsellprice float
        Price for selling a unit to the grid. eg 0.21 for 21 cents per unit (default 0.21)
  -unitsellpriceinflation float
        Rate at which price for selling to the grid increases. Eg for 5% per year use 0.05 here (default 0.04)


```

Example outout:

With default values (4% inflation):

```

Input:
incometaxrate: 0.520000
unitsellprice: 0.210000
unitbuyprice: 0.310000
unitsellpriceinflation: 0.040000
unitbuypriceinflation: 0.040000
solarpanellifetime: 25.000000
selltogridfraction: 0.660000
totalAnnualGeneration: 3200.000000

Year 1:
Sell unit rate (euro): 0.210000
Buy unit rate (euro): 0.310000
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 337.280000
Income from selling this year (euro): 444
Taxable income from selling this year (euro): 244
Income tax due from selling this year (euro): 127
Total income from selling so far (euro): 444
Total savings so far (euro): 337
Total pre-tax income plus savings so far (euro): 781
Total post-tax income plus savings so far (euro): 654
Total income tax due from selling so far (euro): 127

Year 2:
Sell unit rate (euro): 0.218400
Buy unit rate (euro): 0.322400
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 350.771200
Income from selling this year (euro): 461
Taxable income from selling this year (euro): 261
Income tax due from selling this year (euro): 136
Total income from selling so far (euro): 905
Total savings so far (euro): 688
Total pre-tax income plus savings so far (euro): 1593
Total post-tax income plus savings so far (euro): 1330
Total income tax due from selling so far (euro): 262

Year 3:
Sell unit rate (euro): 0.227136
Buy unit rate (euro): 0.335296
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 364.802048
Income from selling this year (euro): 480
Taxable income from selling this year (euro): 280
Income tax due from selling this year (euro): 145
Total income from selling so far (euro): 1384
Total savings so far (euro): 1053
Total pre-tax income plus savings so far (euro): 2437
Total post-tax income plus savings so far (euro): 2029
Total income tax due from selling so far (euro): 408

Year 4:
Sell unit rate (euro): 0.236221
Buy unit rate (euro): 0.348708
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 379.394130
Income from selling this year (euro): 499
Taxable income from selling this year (euro): 299
Income tax due from selling this year (euro): 155
Total income from selling so far (euro): 1883
Total savings so far (euro): 1432
Total pre-tax income plus savings so far (euro): 3316
Total post-tax income plus savings so far (euro): 2752
Total income tax due from selling so far (euro): 563

Year 5:
Sell unit rate (euro): 0.245670
Buy unit rate (euro): 0.362656
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 394.569895
Income from selling this year (euro): 519
Taxable income from selling this year (euro): 319
Income tax due from selling this year (euro): 166
Total income from selling so far (euro): 2402
Total savings so far (euro): 1827
Total pre-tax income plus savings so far (euro): 4229
Total post-tax income plus savings so far (euro): 3500
Total income tax due from selling so far (euro): 729

Year 6:
Sell unit rate (euro): 0.255497
Buy unit rate (euro): 0.377162
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 410.352691
Income from selling this year (euro): 540
Taxable income from selling this year (euro): 340
Income tax due from selling this year (euro): 177
Total income from selling so far (euro): 2942
Total savings so far (euro): 2237
Total pre-tax income plus savings so far (euro): 5179
Total post-tax income plus savings so far (euro): 4273
Total income tax due from selling so far (euro): 906

Year 7:
Sell unit rate (euro): 0.265717
Buy unit rate (euro): 0.392249
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 426.766799
Income from selling this year (euro): 561
Taxable income from selling this year (euro): 361
Income tax due from selling this year (euro): 188
Total income from selling so far (euro): 3503
Total savings so far (euro): 2664
Total pre-tax income plus savings so far (euro): 6167
Total post-tax income plus savings so far (euro): 5073
Total income tax due from selling so far (euro): 1094

Year 8:
Sell unit rate (euro): 0.276346
Buy unit rate (euro): 0.407939
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 443.837471
Income from selling this year (euro): 584
Taxable income from selling this year (euro): 384
Income tax due from selling this year (euro): 199
Total income from selling so far (euro): 4087
Total savings so far (euro): 3108
Total pre-tax income plus savings so far (euro): 7194
Total post-tax income plus savings so far (euro): 5901
Total income tax due from selling so far (euro): 1293

Year 9:
Sell unit rate (euro): 0.287400
Buy unit rate (euro): 0.424256
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 461.590969
Income from selling this year (euro): 607
Taxable income from selling this year (euro): 407
Income tax due from selling this year (euro): 212
Total income from selling so far (euro): 4694
Total savings so far (euro): 3569
Total pre-tax income plus savings so far (euro): 8263
Total post-tax income plus savings so far (euro): 6758
Total income tax due from selling so far (euro): 1505

Year 10:
Sell unit rate (euro): 0.298895
Buy unit rate (euro): 0.441227
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 480.054608
Income from selling this year (euro): 631
Taxable income from selling this year (euro): 431
Income tax due from selling this year (euro): 224
Total income from selling so far (euro): 5325
Total savings so far (euro): 4049
Total pre-tax income plus savings so far (euro): 9374
Total post-tax income plus savings so far (euro): 7645
Total income tax due from selling so far (euro): 1729

Year 11:
Sell unit rate (euro): 0.310851
Buy unit rate (euro): 0.458876
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 499.256792
Income from selling this year (euro): 657
Taxable income from selling this year (euro): 457
Income tax due from selling this year (euro): 237
Total income from selling so far (euro): 5981
Total savings so far (euro): 4549
Total pre-tax income plus savings so far (euro): 10530
Total post-tax income plus savings so far (euro): 8564
Total income tax due from selling so far (euro): 1966

Year 12:
Sell unit rate (euro): 0.323285
Buy unit rate (euro): 0.477231
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 519.227064
Income from selling this year (euro): 683
Taxable income from selling this year (euro): 483
Income tax due from selling this year (euro): 251
Total income from selling so far (euro): 6664
Total savings so far (euro): 5068
Total pre-tax income plus savings so far (euro): 11732
Total post-tax income plus savings so far (euro): 9515
Total income tax due from selling so far (euro): 2217

Year 13:
Sell unit rate (euro): 0.336217
Buy unit rate (euro): 0.496320
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 539.996147
Income from selling this year (euro): 710
Taxable income from selling this year (euro): 510
Income tax due from selling this year (euro): 265
Total income from selling so far (euro): 7374
Total savings so far (euro): 5608
Total pre-tax income plus savings so far (euro): 12982
Total post-tax income plus savings so far (euro): 10500
Total income tax due from selling so far (euro): 2483

Year 14:
Sell unit rate (euro): 0.349665
Buy unit rate (euro): 0.516173
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 561.595993
Income from selling this year (euro): 738
Taxable income from selling this year (euro): 538
Income tax due from selling this year (euro): 280
Total income from selling so far (euro): 8113
Total savings so far (euro): 6169
Total pre-tax income plus savings so far (euro): 14282
Total post-tax income plus savings so far (euro): 11520
Total income tax due from selling so far (euro): 2763

Year 15:
Sell unit rate (euro): 0.363652
Buy unit rate (euro): 0.536820
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 584.059832
Income from selling this year (euro): 768
Taxable income from selling this year (euro): 568
Income tax due from selling this year (euro): 295
Total income from selling so far (euro): 8881
Total savings so far (euro): 6754
Total pre-tax income plus savings so far (euro): 15634
Total post-tax income plus savings so far (euro): 12576
Total income tax due from selling so far (euro): 3058

Year 16:
Sell unit rate (euro): 0.378198
Buy unit rate (euro): 0.558292
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 607.422226
Income from selling this year (euro): 799
Taxable income from selling this year (euro): 599
Income tax due from selling this year (euro): 311
Total income from selling so far (euro): 9680
Total savings so far (euro): 7361
Total pre-tax income plus savings so far (euro): 17041
Total post-tax income plus savings so far (euro): 13671
Total income tax due from selling so far (euro): 3369

Year 17:
Sell unit rate (euro): 0.393326
Buy unit rate (euro): 0.580624
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 631.719115
Income from selling this year (euro): 831
Taxable income from selling this year (euro): 631
Income tax due from selling this year (euro): 328
Total income from selling so far (euro): 10510
Total savings so far (euro): 7993
Total pre-tax income plus savings so far (euro): 18503
Total post-tax income plus savings so far (euro): 14806
Total income tax due from selling so far (euro): 3697

Year 18:
Sell unit rate (euro): 0.409059
Buy unit rate (euro): 0.603849
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 656.987879
Income from selling this year (euro): 864
Taxable income from selling this year (euro): 664
Income tax due from selling this year (euro): 345
Total income from selling so far (euro): 11374
Total savings so far (euro): 8650
Total pre-tax income plus savings so far (euro): 20024
Total post-tax income plus savings so far (euro): 15981
Total income tax due from selling so far (euro): 4043

Year 19:
Sell unit rate (euro): 0.425421
Buy unit rate (euro): 0.628003
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 683.267394
Income from selling this year (euro): 898
Taxable income from selling this year (euro): 698
Income tax due from selling this year (euro): 363
Total income from selling so far (euro): 12273
Total savings so far (euro): 9333
Total pre-tax income plus savings so far (euro): 21606
Total post-tax income plus savings so far (euro): 17200
Total income tax due from selling so far (euro): 4406

Year 20:
Sell unit rate (euro): 0.442438
Buy unit rate (euro): 0.653123
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 710.598090
Income from selling this year (euro): 934
Taxable income from selling this year (euro): 734
Income tax due from selling this year (euro): 382
Total income from selling so far (euro): 13207
Total savings so far (euro): 10044
Total pre-tax income plus savings so far (euro): 23251
Total post-tax income plus savings so far (euro): 18463
Total income tax due from selling so far (euro): 4788

Year 21:
Sell unit rate (euro): 0.460136
Buy unit rate (euro): 0.679248
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 739.022014
Income from selling this year (euro): 972
Taxable income from selling this year (euro): 772
Income tax due from selling this year (euro): 401
Total income from selling so far (euro): 14179
Total savings so far (euro): 10783
Total pre-tax income plus savings so far (euro): 24962
Total post-tax income plus savings so far (euro): 19772
Total income tax due from selling so far (euro): 5189

Year 22:
Sell unit rate (euro): 0.478541
Buy unit rate (euro): 0.706418
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 768.582894
Income from selling this year (euro): 1011
Taxable income from selling this year (euro): 811
Income tax due from selling this year (euro): 422
Total income from selling so far (euro): 15190
Total savings so far (euro): 11551
Total pre-tax income plus savings so far (euro): 26741
Total post-tax income plus savings so far (euro): 21130
Total income tax due from selling so far (euro): 5611

Year 23:
Sell unit rate (euro): 0.497683
Buy unit rate (euro): 0.734675
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 799.326210
Income from selling this year (euro): 1051
Taxable income from selling this year (euro): 851
Income tax due from selling this year (euro): 443
Total income from selling so far (euro): 16241
Total savings so far (euro): 12350
Total pre-tax income plus savings so far (euro): 28591
Total post-tax income plus savings so far (euro): 22538
Total income tax due from selling so far (euro): 6053

Year 24:
Sell unit rate (euro): 0.517590
Buy unit rate (euro): 0.764062
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 831.299258
Income from selling this year (euro): 1093
Taxable income from selling this year (euro): 893
Income tax due from selling this year (euro): 464
Total income from selling so far (euro): 17334
Total savings so far (euro): 13182
Total pre-tax income plus savings so far (euro): 30516
Total post-tax income plus savings so far (euro): 23998
Total income tax due from selling so far (euro): 6518

Year 25:
Sell unit rate (euro): 0.538294
Buy unit rate (euro): 0.794624
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 864.551229
Income from selling this year (euro): 1137
Taxable income from selling this year (euro): 937
Income tax due from selling this year (euro): 487
Total income from selling so far (euro): 18471
Total savings so far (euro): 14046
Total pre-tax income plus savings so far (euro): 32517
Total post-tax income plus savings so far (euro): 25512
Total income tax due from selling so far (euro): 7005


```
