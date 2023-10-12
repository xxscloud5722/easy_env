import Fetch from '@/apis/api.ts';

class KeyApi extends Fetch {
  async list<T>(params: {}) {
    return this.pageTable<T>(await this.post<any>('/', undefined, this.pageParams(params)));
  }
}

export default new KeyApi();
