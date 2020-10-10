// Package transport provides implementation of QUIC transport protocol.
package transport

import (
	"crypto/tls"
	"time"
)

const (
	// ProtocolVersion is the supported QUIC version
	ProtocolVersion = 0xff000000 + 32
	// minProtocolVersion is the minimum version supported.
	minProtocolVersion = 0xff000000 + 29

	// MaxCIDLength is the maximum length of a Connection ID
	MaxCIDLength = 20

	// https://quicwg.org/base-drafts/draft-ietf-quic-transport.html#name-packet-size

	// MaxIPv6PacketSize is the QUIC maximum packet size for IPv6 when Path MTU Discovery is missing.
	MaxIPv6PacketSize = 1232
	// MaxIPv4PacketSize is the QUIC maximum packet size for IPv4 when Path MTU Discovery is missing.
	MaxIPv4PacketSize = 1252
	// MaxPacketSize is the maximum permitted UDP payload.
	MaxPacketSize = 65527
	// MinInitialPacketSize is the QUIC minimum packet size when it contains Initial packet.
	MinInitialPacketSize = 1200
)

// Config is a QUIC connection configuration.
// This implementaton utilizes tls.Config.Rand and tls.Config.Time if available.
type Config struct {
	Version uint32
	TLS     *tls.Config
	Params  Parameters
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	return &Config{
		Version: ProtocolVersion,
		Params: Parameters{
			MaxIdleTimeout:    30 * time.Second,
			MaxUDPPayloadSize: MaxPacketSize,

			InitialMaxData:                 16384,
			InitialMaxStreamDataBidiLocal:  8192,
			InitialMaxStreamDataBidiRemote: 8192,
			InitialMaxStreamDataUni:        8192,
			InitialMaxStreamsBidi:          1,
			InitialMaxStreamsUni:           1,

			AckDelayExponent: 3,
			MaxAckDelay:      25 * time.Millisecond,

			ActiveConnectionIDLimit: 2,
		},
	}
}

// VersionSupported returns true when the provided QUIC transport version is supported.
func VersionSupported(version uint32) bool {
	return version <= ProtocolVersion && version >= minProtocolVersion
}
