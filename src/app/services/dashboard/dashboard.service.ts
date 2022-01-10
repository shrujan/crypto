import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';

export interface User {
  UserName: string,
  Email: string
} 

export interface Coin {
  IsFav?: boolean,
  Name?:  string,
  Symbol?: string
}

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  baseUrl                = 'http://localhost:8081';
  coinList$              = new BehaviorSubject<any>(undefined);
  internationalListInr$  = new BehaviorSubject<any>(undefined);
  purchaseInfo$          = new BehaviorSubject<any>(undefined);
  wazirxData$            = new BehaviorSubject<any>(undefined);
  favoriteCoins$         = new BehaviorSubject<any>(undefined);

  constructor(private http: HttpClient) { }

  getInternationalData(): void {
    let url = `${ this.baseUrl }/getMarketInfo`;
    this.http.get(url).subscribe(coinList => {
      this.internationalListInr$.next(coinList)
    });
  }

  getUsers() {
    let url = `${ this.baseUrl }/getUsers`;
    return this.http.get(url);
  }

  getCoins() {
    let url = `${ this.baseUrl }/getCoin`;
    return this.http.get(url).subscribe((coins: any) => {
      this.favoriteCoins$.next(coins.filter(coin => coin.IsFav));
      this.coinList$.next(coins);
    });
  }

  saveCryptoInfo(data: any) {
    let url = `${ this.baseUrl }/savePurchaseInfo`;
    return this.http.post(url, data);
  }

  sendMail() {
    let url = `${ this.baseUrl }/getMarketInfows`;
    return this.http.get(url);
  }

  updateFavorites(coinData) {
    console.log(coinData)
    let url = `${ this.baseUrl }/favoriteCoin`;
    this.http.post(url, coinData).subscribe((resp) => {
      this.getCoins();
    }, err => this.getCoins());
  }

  fetchAPIWazirx() {
    const url = `${ this.baseUrl }/getMarketInfoWX`;

    this.http.get(url).subscribe(data => {
      this.wazirxData$.next(data)
    })
  }

  getPurchaseInfo(name: string) {
    const url = `${ this.baseUrl }/getPurchaseInfo/${ name }`;

    this.http.get(url).subscribe(purchases => {
      this.purchaseInfo$.next(purchases)
    })
  }

  processPurchasedInfo(purchases: any, wazirData: any, internationalData: any) {
    const uniqueCoins  = [ ... new Set(purchases.map((purchasedCoin: any) => purchasedCoin.coinName)) ]
    const computeQuantity      = (prev: any, current: any) => parseInt(prev.quantity) + parseInt(current.quantity);
    const computePurchasePrice = (prev: any, current: any) => parseInt(prev.purchasePrice) + parseInt(current.purchasePrice);
    const computeTotalAmount   = (prev: any, current: any) => parseInt(prev.totalAmount) + parseInt(current.totalAmount);
    const filterCoins = (coin: any) => purchases.filter((purchasedCoin: any) => purchasedCoin.coinName === coin)
    
    let purchasesArr = [];
    uniqueCoins.forEach(coin => {
      let filteredCoinCount = filterCoins(coin).length;
      let coinObj           = filterCoins(coin)[0];
      coinObj.quantity      = (filteredCoinCount > 1) ? filterCoins(coin).reduce(computeQuantity) : parseInt(filterCoins(coin)[0].quantity);
      coinObj.purchasePrice = (filteredCoinCount > 1) ? filterCoins(coin).reduce(computePurchasePrice) : parseInt(filterCoins(coin)[0].purchasePrice);
      coinObj.totalAmount   = (filteredCoinCount > 1) ? filterCoins(coin).reduce(computeTotalAmount) : parseInt(filterCoins(coin)[0].totalAmount);
      coinObj.indian        = wazirData[coin + 'inr'].buy
      coinObj.international = internationalData.find(data => data.symbol === coin).current_price;
      purchasesArr.push(coinObj)
    });
    return purchasesArr;
  }
}
