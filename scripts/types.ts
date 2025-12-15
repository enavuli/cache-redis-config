// types.ts
export enum RedisConfigSource {
  ENV = 'ENV',
  FILE = 'FILE',
  DEFAULT = 'DEFAULT',
}

export type RedisConfig = {
  host: string;
  port: number;
  db: number;
  password: string;
  prefix: string;
};

export type RedisConfigOptions = {
  host?: string;
  port?: number;
  db?: number;
  password?: string;
  prefix?: string;
  source: RedisConfigSource;
};

export type RedisConfigProvider = {
  getRedisConfig: () => RedisConfig;
};