# solartax
Back of the envelope estimate of income tax due on solar electricity generation

Usage:

```
solartax -h
Usage of ./solartax:
  -annualtaxallowance float
        Income tax allowance. Eg 200.0 for first 200 euro income every year to be untaxed (default 200)
  -incometaxrate float
        Income tax rate. For 40% use 0.4 here. (default 0.4)
  -selltogridfraction float
        Fraction of generated energy sold to grid. Eg if you export 2/3 of generated electricity and use 1/3, set to 0.66. If you use 2/3 and export 1/3, set to 0.33 (default 0.66)
  -solarpanellifetime float
        Lifespan of solar panels (duration of investment). Eg 25.0 for 25 years (default 25)
  -totalannualgeneration float
        Total amount of electricity generated in a year in kilowatts. Eg for 3.2 megawatts set to 3200 (default 3200)
  -unitbuyprice float
        Price for selling a unit to the grid. eg 0.31 for 31 cents per unit (default 0.31)
  -unitbuypriceinflation float
        Rate at which price for selling to the grid increases. Eg for 5% per year use 0.05 here (default 0.04)
  -unitsellprice float
        Price for selling a unit to the grid. eg 0.14 for 14 cents per unit (default 0.14)
  -unitsellpriceinflation float
        Rate at which price for selling to the grid increases. Eg for 5% per year use 0.05 here (default 0.04)

```

Example outout:

With default values (4% inflation):

```

Input:
incometaxrate: 0.400000
unitsellprice: 0.140000
unitbuyprice: 0.310000
unitsellpriceinflation: 0.040000
unitbuypriceinflation: 0.040000
solarpanellifetime: 25.000000
selltogridfraction: 0.660000
totalAnnualGeneration: 3200.000000
Year 1:
Sell unit rate (euro): 0.140000
Buy unit rate (euro): 0.310000
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 337.280000
Income from selling this year (euro): 296
Taxable income from selling this year (euro): 96
Income tax due from selling this year (euro): 38
Total income from selling so far (euro): 296
Total savings so far (euro): 337
Total pre-tax income plus savings so far (euro): 633
Total post-tax income plus savings so far (euro): 595
Total income tax due from selling so far (euro): 38

Year 2:
Sell unit rate (euro): 0.145600
Buy unit rate (euro): 0.322400
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 350.771200
Income from selling this year (euro): 308
Taxable income from selling this year (euro): 108
Income tax due from selling this year (euro): 43
Total income from selling so far (euro): 603
Total savings so far (euro): 688
Total pre-tax income plus savings so far (euro): 1291
Total post-tax income plus savings so far (euro): 1210
Total income tax due from selling so far (euro): 81

Year 3:
Sell unit rate (euro): 0.151424
Buy unit rate (euro): 0.335296
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 364.802048
Income from selling this year (euro): 320
Taxable income from selling this year (euro): 120
Income tax due from selling this year (euro): 48
Total income from selling so far (euro): 923
Total savings so far (euro): 1053
Total pre-tax income plus savings so far (euro): 1976
Total post-tax income plus savings so far (euro): 1847
Total income tax due from selling so far (euro): 129

Year 4:
Sell unit rate (euro): 0.157481
Buy unit rate (euro): 0.348708
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 379.394130
Income from selling this year (euro): 333
Taxable income from selling this year (euro): 133
Income tax due from selling this year (euro): 53
Total income from selling so far (euro): 1256
Total savings so far (euro): 1432
Total pre-tax income plus savings so far (euro): 2688
Total post-tax income plus savings so far (euro): 2506
Total income tax due from selling so far (euro): 182

Year 5:
Sell unit rate (euro): 0.163780
Buy unit rate (euro): 0.362656
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 394.569895
Income from selling this year (euro): 346
Taxable income from selling this year (euro): 146
Income tax due from selling this year (euro): 58
Total income from selling so far (euro): 1601
Total savings so far (euro): 1827
Total pre-tax income plus savings so far (euro): 3428
Total post-tax income plus savings so far (euro): 3188
Total income tax due from selling so far (euro): 241

Year 6:
Sell unit rate (euro): 0.170331
Buy unit rate (euro): 0.377162
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 410.352691
Income from selling this year (euro): 360
Taxable income from selling this year (euro): 160
Income tax due from selling this year (euro): 64
Total income from selling so far (euro): 1961
Total savings so far (euro): 2237
Total pre-tax income plus savings so far (euro): 4198
Total post-tax income plus savings so far (euro): 3894
Total income tax due from selling so far (euro): 304

Year 7:
Sell unit rate (euro): 0.177145
Buy unit rate (euro): 0.392249
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 426.766799
Income from selling this year (euro): 374
Taxable income from selling this year (euro): 174
Income tax due from selling this year (euro): 70
Total income from selling so far (euro): 2335
Total savings so far (euro): 2664
Total pre-tax income plus savings so far (euro): 4999
Total post-tax income plus savings so far (euro): 4625
Total income tax due from selling so far (euro): 374

Year 8:
Sell unit rate (euro): 0.184230
Buy unit rate (euro): 0.407939
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 443.837471
Income from selling this year (euro): 389
Taxable income from selling this year (euro): 189
Income tax due from selling this year (euro): 76
Total income from selling so far (euro): 2724
Total savings so far (euro): 3108
Total pre-tax income plus savings so far (euro): 5832
Total post-tax income plus savings so far (euro): 5382
Total income tax due from selling so far (euro): 450

Year 9:
Sell unit rate (euro): 0.191600
Buy unit rate (euro): 0.424256
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 461.590969
Income from selling this year (euro): 405
Taxable income from selling this year (euro): 205
Income tax due from selling this year (euro): 82
Total income from selling so far (euro): 3129
Total savings so far (euro): 3569
Total pre-tax income plus savings so far (euro): 6698
Total post-tax income plus savings so far (euro): 6167
Total income tax due from selling so far (euro): 532

Year 10:
Sell unit rate (euro): 0.199264
Buy unit rate (euro): 0.441227
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 480.054608
Income from selling this year (euro): 421
Taxable income from selling this year (euro): 221
Income tax due from selling this year (euro): 88
Total income from selling so far (euro): 3550
Total savings so far (euro): 4049
Total pre-tax income plus savings so far (euro): 7599
Total post-tax income plus savings so far (euro): 6979
Total income tax due from selling so far (euro): 620

Year 11:
Sell unit rate (euro): 0.207234
Buy unit rate (euro): 0.458876
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 499.256792
Income from selling this year (euro): 438
Taxable income from selling this year (euro): 238
Income tax due from selling this year (euro): 95
Total income from selling so far (euro): 3988
Total savings so far (euro): 4549
Total pre-tax income plus savings so far (euro): 8536
Total post-tax income plus savings so far (euro): 7821
Total income tax due from selling so far (euro): 715

Year 12:
Sell unit rate (euro): 0.215524
Buy unit rate (euro): 0.477231
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 519.227064
Income from selling this year (euro): 455
Taxable income from selling this year (euro): 255
Income tax due from selling this year (euro): 102
Total income from selling so far (euro): 4443
Total savings so far (euro): 5068
Total pre-tax income plus savings so far (euro): 9511
Total post-tax income plus savings so far (euro): 8694
Total income tax due from selling so far (euro): 817

Year 13:
Sell unit rate (euro): 0.224145
Buy unit rate (euro): 0.496320
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 539.996147
Income from selling this year (euro): 473
Taxable income from selling this year (euro): 273
Income tax due from selling this year (euro): 109
Total income from selling so far (euro): 4916
Total savings so far (euro): 5608
Total pre-tax income plus savings so far (euro): 10524
Total post-tax income plus savings so far (euro): 9598
Total income tax due from selling so far (euro): 926

Year 14:
Sell unit rate (euro): 0.233110
Buy unit rate (euro): 0.516173
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 561.595993
Income from selling this year (euro): 492
Taxable income from selling this year (euro): 292
Income tax due from selling this year (euro): 117
Total income from selling so far (euro): 5409
Total savings so far (euro): 6169
Total pre-tax income plus savings so far (euro): 11578
Total post-tax income plus savings so far (euro): 10535
Total income tax due from selling so far (euro): 1043

Year 15:
Sell unit rate (euro): 0.242435
Buy unit rate (euro): 0.536820
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 584.059832
Income from selling this year (euro): 512
Taxable income from selling this year (euro): 312
Income tax due from selling this year (euro): 125
Total income from selling so far (euro): 5921
Total savings so far (euro): 6754
Total pre-tax income plus savings so far (euro): 12674
Total post-tax income plus savings so far (euro): 11506
Total income tax due from selling so far (euro): 1168

Year 16:
Sell unit rate (euro): 0.252132
Buy unit rate (euro): 0.558292
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 607.422226
Income from selling this year (euro): 533
Taxable income from selling this year (euro): 333
Income tax due from selling this year (euro): 133
Total income from selling so far (euro): 6453
Total savings so far (euro): 7361
Total pre-tax income plus savings so far (euro): 13814
Total post-tax income plus savings so far (euro): 12513
Total income tax due from selling so far (euro): 1301

Year 17:
Sell unit rate (euro): 0.262217
Buy unit rate (euro): 0.580624
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 631.719115
Income from selling this year (euro): 554
Taxable income from selling this year (euro): 354
Income tax due from selling this year (euro): 142
Total income from selling so far (euro): 7007
Total savings so far (euro): 7993
Total pre-tax income plus savings so far (euro): 15000
Total post-tax income plus savings so far (euro): 13557
Total income tax due from selling so far (euro): 1443

Year 18:
Sell unit rate (euro): 0.272706
Buy unit rate (euro): 0.603849
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 656.987879
Income from selling this year (euro): 576
Taxable income from selling this year (euro): 376
Income tax due from selling this year (euro): 150
Total income from selling so far (euro): 7583
Total savings so far (euro): 8650
Total pre-tax income plus savings so far (euro): 16233
Total post-tax income plus savings so far (euro): 14639
Total income tax due from selling so far (euro): 1593

Year 19:
Sell unit rate (euro): 0.283614
Buy unit rate (euro): 0.628003
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 683.267394
Income from selling this year (euro): 599
Taxable income from selling this year (euro): 399
Income tax due from selling this year (euro): 160
Total income from selling so far (euro): 8182
Total savings so far (euro): 9333
Total pre-tax income plus savings so far (euro): 17515
Total post-tax income plus savings so far (euro): 15762
Total income tax due from selling so far (euro): 1753

Year 20:
Sell unit rate (euro): 0.294959
Buy unit rate (euro): 0.653123
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 710.598090
Income from selling this year (euro): 623
Taxable income from selling this year (euro): 423
Income tax due from selling this year (euro): 169
Total income from selling so far (euro): 8805
Total savings so far (euro): 10044
Total pre-tax income plus savings so far (euro): 18848
Total post-tax income plus savings so far (euro): 16926
Total income tax due from selling so far (euro): 1922

Year 21:
Sell unit rate (euro): 0.306757
Buy unit rate (euro): 0.679248
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 739.022014
Income from selling this year (euro): 648
Taxable income from selling this year (euro): 448
Income tax due from selling this year (euro): 179
Total income from selling so far (euro): 9453
Total savings so far (euro): 10783
Total pre-tax income plus savings so far (euro): 20235
Total post-tax income plus savings so far (euro): 18134
Total income tax due from selling so far (euro): 2101

Year 22:
Sell unit rate (euro): 0.319028
Buy unit rate (euro): 0.706418
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 768.582894
Income from selling this year (euro): 674
Taxable income from selling this year (euro): 474
Income tax due from selling this year (euro): 190
Total income from selling so far (euro): 10126
Total savings so far (euro): 11551
Total pre-tax income plus savings so far (euro): 21678
Total post-tax income plus savings so far (euro): 19387
Total income tax due from selling so far (euro): 2291

Year 23:
Sell unit rate (euro): 0.331789
Buy unit rate (euro): 0.734675
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 799.326210
Income from selling this year (euro): 701
Taxable income from selling this year (euro): 501
Income tax due from selling this year (euro): 200
Total income from selling so far (euro): 10827
Total savings so far (euro): 12350
Total pre-tax income plus savings so far (euro): 23178
Total post-tax income plus savings so far (euro): 20687
Total income tax due from selling so far (euro): 2491

Year 24:
Sell unit rate (euro): 0.345060
Buy unit rate (euro): 0.764062
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 831.299258
Income from selling this year (euro): 729
Taxable income from selling this year (euro): 529
Income tax due from selling this year (euro): 212
Total income from selling so far (euro): 11556
Total savings so far (euro): 13182
Total pre-tax income plus savings so far (euro): 24738
Total post-tax income plus savings so far (euro): 22035
Total income tax due from selling so far (euro): 2702

Year 25:
Sell unit rate (euro): 0.358863
Buy unit rate (euro): 0.794624
Generated units (kilowatts): 3200
Value of consumed solar electricity units for this year (euro): 864.551229
Income from selling this year (euro): 758
Taxable income from selling this year (euro): 558
Income tax due from selling this year (euro): 223
Total income from selling so far (euro): 12314
Total savings so far (euro): 14046
Total pre-tax income plus savings so far (euro): 26360
Total post-tax income plus savings so far (euro): 23435
Total income tax due from selling so far (euro): 2926

```
