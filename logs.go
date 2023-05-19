package ctstream

// A list of logs used by Chrome.
// From: https://source.chromium.org/chromium/chromium/src/+/main:components/certificate_transparency/data/log_list.json
var Logs = []Log{
	{Name: "Argon2022", URI: "https://ct.googleapis.com/logs/argon2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeIPc6fGmuBg6AJkv/z7NFckmHvf/OqmjchZJ6wm2qN200keRDg352dWpi7CHnSV51BpQYAj1CQY5JuRAwrrDwg=="},
	{Name: "Argon2023", URI: "https://ct.googleapis.com/logs/argon2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0JCPZFJOQqyEti5M8j13ALN3CAVHqkVM4yyOcKWCu2yye5yYeqDpEXYoALIgtM3TmHtNlifmt+4iatGwLpF3eA=="},
	{Name: "Argon2024", URI: "https://ct.googleapis.com/logs/us1/argon2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA=="},
	{Name: "Xenon2022", URI: "https://ct.googleapis.com/logs/xenon2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+WS9FSxAYlCVEzg8xyGwOrmPonoV14nWjjETAIdZvLvukPzIWBMKv6tDNlQjpIHNrUcUt1igRPpqoKDXw2MeKw=="},
	{Name: "Xenon2023", URI: "https://ct.googleapis.com/logs/xenon2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEchY+C+/vzj5g3ZXLY3q5qY1Kb2zcYYCmRV4vg6yU84WI0KV00HuO/8XuQqLwLZPjwtCymeLhQunSxgAnaXSuzg=="},
	{Name: "Xenon2024", URI: "https://ct.googleapis.com/logs/eu1/xenon2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEuWDgNB415GUAk0+QCb1a7ETdjA/O7RE+KllGmjG2x5n33O89zY+GwjWlPtwpurvyVOKoDIMIUQbeIW02UI44TQ=="},
	{Name: "Nimbus2022", URI: "https://ct.cloudflare.com/logs/nimbus2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESLJHTlAycmJKDQxIv60pZG8g33lSYxYpCi5gteI6HLevWbFVCdtZx+m9b+0LrwWWl/87mkNN6xE0M4rnrIPA/w=="},
	{Name: "Nimbus2023", URI: "https://ct.cloudflare.com/logs/nimbus2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEi/8tkhjLRp0SXrlZdTzNkTd6HqmcmXiDJz3fAdWLgOhjmv4mohvRhwXul9bgW0ODgRwC9UGAgH/vpGHPvIS1qA=="},
	{Name: "Nimbus2024", URI: "https://ct.cloudflare.com/logs/nimbus2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEd7Gbe4/mizX+OpIpLayKjVGKJfyTttegiyk3cR0zyswz6ii5H+Ksw6ld3Ze+9p6UJd02gdHrXSnDK0TxW8oVSA=="},
	{Name: "Yeti2023", URI: "https://yeti2023.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfQ0DsdWYitzwFTvG3F4Nbj8Nv5XIVYzQpkyWsU4nuSYlmcwrAp6m092fsdXEw6w1BAeHlzaqrSgNfyvZaJ9y0Q=="},
	{Name: "Yeti2024", URI: "https://yeti2024.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV7jBbzCkfy7k8NDZYGITleN6405Tw7O4c4XBGA0jDliE0njvm7MeLBrewY+BGxlEWLcAd2AgGnLYgt6unrHGSw=="},
	{Name: "Yeti2025", URI: "https://yeti2025.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE35UAXhDBAfc34xB00f+yypDtMplfDDn+odETEazRs3OTIMITPEy1elKGhj3jlSR82JGYSDvw8N8h8bCBWlklQw=="},
	{Name: "Nessie2023", URI: "https://nessie2023.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEXu8iQwSCRSf2CbITGpUpBtFVt8+I0IU0d1C36Lfe1+fbwdaI0Z5FktfM2fBoI1bXBd18k2ggKGYGgdZBgLKTg=="},
	{Name: "Nessie2024", URI: "https://nessie2024.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAELfyieza/VpHp/j/oPfzDp+BhUuos6QWjnycXgQVwa4FhRIr4OxCAQu0DLwBQIfxBVISjVNUusnoWSyofK2YEKw=="},
	{Name: "Nessie2025", URI: "https://nessie2025.ct.digicert.com/log/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8vDwp4uBLgk5O59C2jhEX7TM7Ta72EN/FklXhwR/pQE09+hoP7d4H2BmLWeadYC3U6eF1byrRwZV27XfiKFvOA=="},
	{Name: "Sabre", URI: "https://sabre.ct.comodo.com/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8m/SiQ8/xfiHHqtls9m7FyOMBg4JVZY9CgiixXGz0akvKD6DEL8S0ERmFe9U4ZiA0M4kbT5nmuk3I85Sk4bagA=="},
	{Name: "Mammoth", URI: "https://mammoth.ct.comodo.com/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7+R9dC4VFbbpuyOL+yy14ceAmEf7QGlo/EmtYU6DRzwat43f/3swtLr/L8ugFOOt1YU/RFmMjGCL17ixv66MZw=="},
	{Name: "Oak2022", URI: "https://oak.ct.letsencrypt.org/2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhjyxDVIjWt5u9sB/o2S8rcGJ2pdZTGA8+IpXhI/tvKBjElGE5r3de4yAfeOPhqTqqc+o7vPgXnDgu/a9/B+RLg=="},
	{Name: "Oak2023", URI: "https://oak.ct.letsencrypt.org/2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsz0OeL7jrVxEXJu+o4QWQYLKyokXHiPOOKVUL3/TNFFquVzDSer7kZ3gijxzBp98ZTgRgMSaWgCmZ8OD74mFUQ=="},
	{Name: "Oak2024H1", URI: "https://oak.ct.letsencrypt.org/2024h1/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEVkPXfnvUcre6qVG9NpO36bWSD+pet0Wjkv3JpTyArBog7yUvuOEg96g6LgeN5uuk4n0kY59Gv5RzUo2Wrqkm/Q=="},
	{Name: "Oak2024H2", URI: "https://oak.ct.letsencrypt.org/2024h2/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE13PWU0fp88nVfBbC1o9wZfryUTapE4Av7fmU01qL6E8zz8PTidRfWmaJuiAfccvKu5+f81wtHqOBWa+Ss20waA=="},
	{Name: "TrustAsia2022", URI: "https://ct.trustasia.com/log2022/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEu1LyFs+SC8555lRtwjdTpPX5OqmzBewdvRbsMKwu+HliNRWOGtgWLuRIa/bGE/GWLlwQ/hkeqBi4Dy3DpIZRlw=="},
	{Name: "TrustAsia2023", URI: "https://ct.trustasia.com/log2023/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpBFS2xdBTpDUVlESMFL4mwPPTJ/4Lji18Vq6+ji50o8agdqVzDPsIShmxlY+YDYhINnUrF36XBmhBX3+ICP89Q=="},
	{Name: "TrustAsia2024-2", URI: "https://ct2024.trustasia.com/log2024/", PubKey: "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEp2TieYE/YdfsxvhlKB2gtGYzwyXVCpV4nI/+pCrYj35y4P6of/ixLYXAjhJ0DS+Mq9d/eh7ZhDM56P2JX5ZICA=="},
}

// FindByName returns the Log with the given name from Logs.
// If Log not found, returns nil.
func FindByName(name string) *Log {

	for i := range Logs {
		if Logs[i].Name == name {
			return &Logs[i]
		}
	}

	return nil
}

// FindByURI returns the Log with the given URI from Logs.
// If Log not found, returns nil.
func FindByURI(uri string) *Log {

	for i := range Logs {
		if Logs[i].URI == uri {
			return &Logs[i]
		}
	}

	return nil
}

// FindByPubKey returns the Log with the given public key from Logs.
// If Log not found, returns nil.
func FindByPubKey(key string) *Log {

	for i := range Logs {
		if Logs[i].PubKey == key {
			return &Logs[i]
		}
	}

	return nil
}
