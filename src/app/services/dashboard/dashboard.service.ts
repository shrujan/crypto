import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';

export interface User {
  UserName: string,
  Email: string
} 

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  // baseUrl = 'https://api.coingecko.com/api/v3';
  baseUrl        = 'http://localhost:8081';
  coinListInr$  = new BehaviorSubject<any>(undefined);
  purchaseInfo$ = new BehaviorSubject<any>(undefined) 
  wazirxData$   = new BehaviorSubject<any>(undefined);


  constructor(private http: HttpClient) { }

  getCoinsInfoInr(): void {
    let url = `${ this.baseUrl }/getMarketInfo`;
    this.http.get(url).subscribe(coinList => {
      this.coinListInr$.next(coinList)
    });
  }

  getUsers() {
    let url = `${ this.baseUrl }/getUsers`;
    return this.http.get(url);
  }

  saveCryptoInfo(data: any) {
    let url = `${ this.baseUrl }/savePurchaseInfo`;
    return this.http.post(url, data);
  }

  sendMail() {
    let url = `${ this.baseUrl }/getMarketInfows`;
    return this.http.get(url);
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

  processPurchasedInfo(purchases: any, coinData: any) {
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
      purchasesArr.push(coinObj)
    });
    return purchasesArr;
  }
}
