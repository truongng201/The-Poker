import React from "react";
import "./TablePoker.css";

export default function TablePoker() {
  return (
    <div class="TablePoker">
      <div className="card-place">
        <div className="card">A</div>
        <div className="card">2</div>
        <div className="card">3</div>
        <div className="card">4</div>
        <div className="card">5</div>
      </div>
      <div className="players">
        <div className="player player-1">
          <div className="bank bank-up">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=1"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-2">
          <div className="bank bank-down">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=2"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-3">
          <div className="bank bank-down">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=3"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-4">
          <div className="bank bank-down">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=4"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-5">
          <div className="bank bank-down">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=5"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-6">
          <div className="bank bank-up">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=6"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-7">
          <div className="bank bank-up">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=7"
          />
          <div className="player-name"></div>
        </div>
        <div className="player player-8">
          <div className="bank bank-up">
            <div className="bank-value">1000000</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=8"
          />
          <div className="player-name"></div>
        </div>
      </div>
    </div>
  );
}
