import React from "react";
import "./TablePoker.css";
import Card from "./Card";

export default function TablePoker() {
  return (
    <div class="TablePoker">
      <div className="card-place">
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
      </div>
      <div className="players">
        <div className="player player-1 ">
          <div className="bank bank-up">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=1"
          />
          <div className="player-group-cards player-group-cards-right">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-right">Messi</div>
        </div>
        <div className="player player-2">
          <div className="bank bank-down">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=2"
          />
          <div className="player-group-cards player-group-cards-down">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-down">John</div>
        </div>
        <div className="player player-3">
          <div className="bank bank-down">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=3"
          />
          <div className="player-group-cards player-group-cards-down">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-down">Mia</div>
        </div>
        <div className="player player-4">
          <div className="bank bank-down">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=4"
          />
          <div className="player-group-cards player-group-cards-down">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-down">Terry</div>
        </div>
        <div className="player player-5">
          <div className="bank bank-down">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=5"
          />
          <div className="player-group-cards player-group-cards-left">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-left">Halland</div>
        </div>
        <div className="player player-6">
          <div className="bank bank-up">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=6"
          />
          <div className="player-group-cards player-group-cards-up">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-up">Mary</div>
        </div>
        <div className="player player-7">
          <div className="bank bank-up">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=7"
          />
          <div className="player-group-cards player-group-cards-up">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-up">Tony</div>
        </div>
        <div className="player player-8">
          <div className="bank bank-up">
            <div className="bank-value">1.000.000$</div>
            <div className="jetons v-10"></div>
            <div className="jetons v-5"></div>
            <div className="jetons v-1"></div>
          </div>
          <img
            className="avatar"
            alt="avatar"
            src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=8"
          />
          <div className="player-group-cards player-group-cards-up">
            <Card />
            <Card />
          </div>
          <div className="player-name player-name-up">Henry</div>
        </div>
      </div>
    </div>
  );
}
