package url

import (
	"math"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/mutagen-io/mutagen/pkg/url/forwarding"
)

// Supported returns whether or not a URL kind is supported.
func (k Kind) Supported() bool {
	switch k {
	case Kind_Synchronization:
		return true
	case Kind_Forwarding:
		return true
	default:
		return false
	}
}

// MightRequireInput indicates whether or not a protocol might require
// interactive user input (e.g. a password) while connecting.
func (p Protocol) MightRequireInput() bool {
	switch p {
	case Protocol_SSH:
		return true
	default:
		return false
	}
}

// EnsureValid ensures that URL's invariants are respected.
func (u *URL) EnsureValid() error {
	// Ensure that the URL is non-nil.
	if u == nil {
		return errors.New("nil URL")
	}

	// Ensure that the kind is supported.
	if !u.Kind.Supported() {
		return errors.New("unsupported URL kind")
	}

	// Validate the User, Host, Port, and Environment components based on
	// protocol.
	if u.Protocol == Protocol_Local {
		if u.User != "" {
			return errors.New("local URL with non-empty username")
		} else if u.Host != "" {
			return errors.New("local URL with non-empty hostname")
		} else if u.Port != 0 {
			return errors.New("local URL with non-zero port")
		} else if len(u.Environment) != 0 {
			return errors.New("local URL with environment variables")
		}
	} else if u.Protocol == Protocol_SSH {
		if u.Host == "" {
			return errors.New("SSH URL with empty hostname")
		} else if u.Port > math.MaxUint16 {
			return errors.New("SSH URL with invalid port")
		} else if len(u.Environment) != 0 {
			return errors.New("SSH URL with environment variables")
		}
	} else if u.Protocol == Protocol_Tunnel {
		if u.User != "" {
			return errors.New("tunnel URL with non-empty username")
		} else if u.Host == "" {
			return errors.New("tunnel URL with empty tunnel identifier/name")
		} else if u.Port != 0 {
			return errors.New("tunnel URL with non-zero port")
		} else if len(u.Environment) != 0 {
			return errors.New("tunnel URL with environment variables")
		}
	} else if u.Protocol == Protocol_Docker {
		// In the case of Docker, we intentionally avoid validating environment
		// variables since the values used could change over time. Since we
		// default to empty values for unspecified environment variables, this
		// works out fine, at least so long as Docker continues to treat empty
		// environment variables the same as unspecified ones.
		if u.Host == "" {
			return errors.New("Docker URL with empty container identifier")
		} else if u.Port != 0 {
			return errors.New("Docker URL with non-zero port")
		}
	} else {
		return errors.New("unknown or unsupported protocol")
	}

	// Validate the path component depending on the URL kind.
	if u.Kind == Kind_Synchronization {
		// Ensure the path is non-empty.
		if u.Path == "" {
			return errors.New("empty path")
		}

		// If this is a local URL, then ensure that the path is absolute.
		if u.Protocol == Protocol_Local && !filepath.IsAbs(u.Path) {
			return errors.New("local URL with relative path")
		}

		// If this is a tunnel or Docker URL, we can actually do a bit of
		// additional validation.
		if u.Protocol == Protocol_Tunnel || u.Protocol == Protocol_Docker {
			if !(u.Path[0] == '/' || u.Path[0] == '~' || isWindowsPath(u.Path)) {
				return errors.New("incorrect first path character")
			}
		}
	} else if u.Kind == Kind_Forwarding {
		// Parse the forwarding endpoint URL to ensure that it's valid.
		protocol, address, err := forwarding.Parse(u.Path)
		if err != nil {
			return errors.Wrap(err, "invalid forwarding endpoint URL")
		}

		// If this is a local URL and represents a Unix domain socket endpoint,
		// then ensure that the socket path is absolute.
		if u.Protocol == Protocol_Local && protocol == "unix" && !filepath.IsAbs(address) {
			return errors.New("local Unix domain socket URL with relative path")
		}
	}

	// Success.
	return nil
}
