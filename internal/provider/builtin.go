package provider

// Built-in providers shipped with the app. Users can add their own
// Subscription entries on top of these.
//
// On the "Mahsa servers" source: MahsaNG's primary rotating feed from
// mahsaserver.com is delivered over a closed, authenticated, encrypted protocol
// that is deliberately not published (see the project README). We therefore
// cannot reproduce that exact endpoint, and the public collector the app used
// to reference (YeBeKhe) has since been taken down. Instead we ship several
// live public free-config aggregators. If the authenticated mahsaserver
// protocol is later obtained legitimately (e.g. from the maintainers, who run
// the open NikaNG fork), it slots in as just another Provider implementation
// without touching the rest of the app.
//
// NOTE: these are community-run sources whose availability changes over time.
// Multiple are registered so one going offline does not break "GET CONFIG".
//
// weight is how many configs a provider contributes per round in GET CONFIG's
// selection: higher = a larger share of the capped list. Weights favour the
// Iran-oriented sources. Dropped 2026-06: the RU-White whitelist sources (built
// for Russia; field-tested, every sampled one failed from Iran) and the
// mahdibland Aggregator (its feed went empty).
var builtinSources = []struct {
	name, url string
	weight    int
}{
	// Iran-oriented, auto-tested "healthy only", carries REALITY.
	{"Freedom-VLESS", "https://api.github.com/repos/MahanKenway/Freedom-V2Ray/contents/configs/vless.txt?ref=main", 3},
	{"BarryFar-VLESS", "https://api.github.com/repos/barry-far/V2ray-config/contents/Splitted-By-Protocol/vless.txt?ref=main", 3},
	// Filtered subscription, carries REALITY.
	{"MatinGhanbari-VLESS", "https://api.github.com/repos/MatinGhanbari/v2ray-configs/contents/subscriptions/filtered/subs/vless.txt?ref=main", 2},
	{"BarryFar", "https://api.github.com/repos/barry-far/V2ray-Config/contents/All_Configs_Sub.txt?ref=main", 2},
	{"MhdiTaheri", "https://api.github.com/repos/MhdiTaheri/V2rayCollector/contents/sub/mix?ref=main", 2},
}

// serverlessURL is the patterniha/Serverless-for-Iran feed: whole xray configs
// (fragment/noise/routing, no proxy server) that bypass filtering server-side.
// Unlike the scraped sources these never go dead, so they ride along as a
// reliable fallback that's always present in the list.
const serverlessURL = "https://raw.githubusercontent.com/patterniha/Serverless-for-Iran/refs/heads/main/Subscription/Serverless-for-Iran.json"

// Builtins returns the providers enabled by default.
func Builtins() []Provider {
	out := make([]Provider, 0, len(builtinSources)+1)
	for _, s := range builtinSources {
		out = append(out, NewWeightedSubscription(s.name, s.url, s.weight))
	}
	out = append(out, NewServerless("Serverless", serverlessURL))
	return out
}
