import Fetch from 'beer-network/api';

class KeyApi extends Fetch {
  async list<T>(keyword: string) {
    return this.get<T>(`/pair/list${keyword === undefined || keyword === '' ? '' : '/' + keyword}`, undefined);
  }

  async save(params: {}) {
    return this.post<any>('/pair/save', undefined, params);
  }

  async remove(key: string) {
    return this.post<any>('/pair/remove', undefined, { key });
  }
}

export default new KeyApi(import.meta.env.VITE_REQUEST_BASE_URL?.toString());
