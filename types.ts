/**
 * Types and interfaces for the cache-redis-config library
 */

export interface RedisConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
  ttl?: number;
  maxRedirects?: number;
}

export interface Config {
  redis: RedisConfig;
}