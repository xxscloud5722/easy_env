import Fetch from '@/apis/api.ts';

class ScriptApi extends Fetch {
  async list<T>(params: {}) {
    return this.pageTable<T>(await this.get<any>('/script/list', params));
  }

  async directoryList<T>() {
    return this.get<T>('/script-directory/list', undefined);
  }

  async save<T>(params: {}) {
    return this.post<T>('/script/save', undefined, params);
  }

  async remove<T>(id: number) {
    return this.post<T>('/script/remove', undefined, { id });
  }

  async removeDirectory<T>(id: number) {
    return this.post<T>('/script-directory/remove', undefined, { id });
  }

  async createDirectory<T>(params: {}) {
    return this.post<T>('/script-directory/create', undefined, params);
  }

  async renameDirectory<T>(params: {}) {
    return this.post<T>('/script-directory/rename', undefined, params);
  }
}

export default new ScriptApi();
