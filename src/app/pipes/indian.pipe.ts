import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'indian'
})
export class IndianPipe implements PipeTransform {
  transform(value: number, ...args: unknown[]): string {
    value = Number(value)
    return value.toLocaleString('en-IN', {
      maximumFractionDigits: 6,
      style: 'currency',
      currency: 'INR'
    });
  }
}
