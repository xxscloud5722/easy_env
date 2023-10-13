import { RequestData } from '@ant-design/pro-components';
import session from '@/apis/session.ts';

export type JsonResponse<T> = {
  success: boolean;
  data: T;
  code: string;
  message?: string;
};

export default class Fetch {
  private pathPrefix = '/api';

  async post<T>(path: string, params?: any, body?: any): Promise<JsonResponse<T>> {
    const paramQuery = this.searchParams(params);
    const response = await fetch(this.pathPrefix + path + (paramQuery.toString() === '' ? '' : '?') + paramQuery.toString(), {
      method: 'POST',
      headers: {
        ...this.authorization(),
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body || {})
    });
    return this.responseJson<T>(response);
  }

  async get<T>(path: string, params?: any): Promise<JsonResponse<T>> {
    const paramQuery = this.searchParams(params);
    const response = await fetch(this.pathPrefix + path + '?' + paramQuery.toString(), {
      method: 'GET',
      headers: {
        ...this.authorization()
      }
    });
    return this.responseJson<T>(response);
  }

  private searchParams(params?: any): URLSearchParams {
    const paramQuery = new URLSearchParams();
    Object.keys(params || {})
      .forEach(key => {
        paramQuery.append(key, params[key]);
      });
    const defaultQuery: any = { 'access-token': session.getToken() };
    Object.keys(defaultQuery || {})
      .forEach(key => {
        paramQuery.append(key, defaultQuery[key]);
      });
    return paramQuery;
  }

  async cache<T>(key: string, callback: () => Promise<T>) {
    const cacheKey = 'CACHE_' + key;
    const cacheValue = localStorage.getItem(cacheKey);
    if (cacheValue !== null) {
      return JSON.parse(cacheValue) as T;
    }
    const response = await callback();
    localStorage.setItem(cacheKey, JSON.stringify(response));
    return response;
  }

  async cachePage<T>(key: string, callback: () => Promise<T>) {
    await this.sleep(100);
    const cacheKey = 'CACHE_' + key;
    const cacheValue = (window as any)[cacheKey];
    if (cacheValue !== null && cacheValue !== undefined) {
      return JSON.parse(cacheValue);
    }
    const response = await callback();
    (window as any)[cacheKey] = JSON.stringify(response);
    return response as T;
  }

  authorization() {
    return {};
  }

  async responseJson<T>(response: Response): Promise<JsonResponse<T>> {
    return await response.json() as JsonResponse<T>;
  }

  pageParams(params: any) {
    return {
      ...params,
      currentPage: params.current,
      current: undefined
    };
  }

  async pageTable<T>(response: JsonResponse<any>): Promise<Partial<RequestData<T>>> {
    return {
      data: (response.data || {}).records || (response.data || []),
      page: (response.data || {}).currentPage || 1,
      total: (response.data || {}).totalSize || (response.data || []).length,
      success: response.success
    } as RequestData<T>;
  }

  async sleep(value: number) {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(undefined);
      }, value);
    });
  }
}
