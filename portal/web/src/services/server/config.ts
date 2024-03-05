import 'server-only'
import path from 'path'
import fs from 'fs'

interface IServerConfig {
    ENV: string,
    SELF_URL: string,
    SERVER: string,
    RESOURCE_SERVER: string,
    AES_KEY: string,
    AES_IV: string,
    REDIS: string,
    REDIS_PASSWORD: string,
    AUTH_SERVER: string,
    ClientId: string,
    ClientSecret: string
    SIGN: {
        PASSWORD: {
            ENABLE: boolean,
        },
        WEBAUTHN: {
            ENABLE: boolean,
        },
        EMAIL: {
            ENABLE: boolean,
        },
    }
}

let serverConfig: IServerConfig | undefined
 
export async function loadServerConfig (): Promise<IServerConfig> {
  let configPath = process.env.NODE_ENV ? `./runtime/server.${process.env.NODE_ENV}.json` : './runtime/server.json'
    
  if (!path.isAbsolute(configPath)) {
    configPath = path.join(process.cwd(), configPath)
  }
  
  const fileContent = fs.readFileSync(configPath, 'utf8')
    
  serverConfig = JSON.parse(fileContent) as IServerConfig

  return serverConfig
}

