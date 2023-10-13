import Fetch from '@/apis/api.ts';

class KeyApi extends Fetch {
  async list<T>(keyword: string) {
    return this.pageTable<T>(await this.get<any>(`/pair/list${keyword === undefined || keyword === '' ? '' : '/' + keyword}`, undefined));
  }

  async save(params: {}) {
    return this.post<any>('/pair/save', undefined, params);
  }

  async remove(key: string) {
    return this.post<any>('/pair/remove', undefined, { key });
  }
}

export default new KeyApi();
