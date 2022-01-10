import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../services/dashboard/dashboard.service';
import { find, findIndex } from 'lodash'; 

interface FavoriteCoin {
  id: string,
  favorite: boolean,
  purchasedPrice: number,
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  coinList: any;
  coins: any;
  displayedColumns: string[] = ['Favorites', 'Coin', 'Price', '24 High', '24 Low', 'Change %', 'Purchased'];
  favorites: FavoriteCoin[] = [];
  filteredCoins: any;

  constructor(private dashboardService: DashboardService) { }

  ngOnInit(): void {
    this.dashboardService.internationalListInr$.subscribe(coinList => {
      if (!coinList) return
      this.setCustomPrices(coinList);
      coinList.forEach(coin => {
        coin.favorite = this.isFav(coin.symbol)
      });
      this.coins = coinList;
      this.filteredCoins = coinList;
    })

    this.dashboardService.coinList$.subscribe(coins => this.coinList = coins)
    // const source = interval(1000000);
    // source.subscribe(() => this.getCoinData())
    
  }

  savePurchasedPrice(coin: any): void {
    const updatePrice = (favCoin: FavoriteCoin) => favCoin.purchasedPrice = coin.purchasedPrice;
    const favCoin = find(this.favorites, fav => fav.id === coin.id);
    (!favCoin) ? this.saveNewFavCoin(coin) : updatePrice(favCoin);
    this.setCustomPrices(this.filteredCoins)
  }

  setCustomPrices(coinList: any): void {
    coinList.forEach((coin: any) => {
      const favCoin = find(this.favorites, fav => fav.id === coin.id)
      coin.favorite = false;
      coin.purchasedPrice = 0;
      coin.differenceFromPurchase = 0;
      if (favCoin) {
        coin.favorite = true;
        coin.purchasedPrice = favCoin.purchasedPrice,
        coin.differenceFromPurchase = this.calcPercentage(Number(coin.purchasedPrice), Number(coin.current_price))
      }
    });
  }

  setFavorite(coin: any) {
    const deleteFavCoin = () => {
      this.updateFavStatusOfCoins(coin, false);
    }
    // const favCoinIndex = findIndex(this.favorites, (fav: FavoriteCoin) => fav.id === coin.id);
    (!this.isFav(coin.symbol)) ? this.saveNewFavCoin(coin) : deleteFavCoin();
    this.setCustomPrices(this.filteredCoins)
  }

  calcPercentage(purchasedPrice: number, currentPrice: number): number {
    return (!purchasedPrice) ? 0 : (currentPrice - purchasedPrice) / purchasedPrice * 100;
  }

  saveNewFavCoin(coin: any): void {
    // this.favorites.push({
    //   id: coin.id,
    //   favorite: true,
    //   purchasedPrice: coin.purchasedPrice,
    // });
    this.updateFavStatusOfCoins(coin, true);
  }

  updateFavStatusOfCoins(coin: any, isFav: boolean): void {
    const params = {
      isFav:  isFav,
      symbol: coin.symbol
    }
    this.dashboardService.updateFavorites(params);
  }

  isFav(coinSymbol: string): boolean {
    if (!coinSymbol || !this.coinList) return false;
    if (this.coinList.find(coin => coin.Symbol === coinSymbol) === undefined) {
      console.log(coinSymbol)
      return false;
    }
    return this.coinList.find(coin => coin.Symbol === coinSymbol).IsFav;
  }
}
