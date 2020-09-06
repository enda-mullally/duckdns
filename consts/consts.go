package consts

// AppVersion - Replace this value in your build pipeline or use ldflags etc, for now hardcoding
const AppVersion string = "1.0.1"

// UpdateDomainsEnvName - The name of the env var containing the domains value when run in 'update --env' mode
const UpdateDomainsEnvName = "DUCKDNS_DOMAINS"

// UpdateTokenEnvName - The name of the env var containing the token value when run in 'update --env' mode
const UpdateTokenEnvName = "DUCKDNS_TOKEN"

// UpdateIPEnvName - The name of the env var containing the ip value when run in 'update --env' mode
const UpdateIPEnvName = "DUCKDNS_IP"

// DuckDNSOrg - The official name of the service this cli is built for
const DuckDNSOrg = "DuckDNS.org"
