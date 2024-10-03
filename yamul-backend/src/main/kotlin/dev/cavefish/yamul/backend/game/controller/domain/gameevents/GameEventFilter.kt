package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.game.controller.domain.GameState

interface GameEventFilter {
    operator fun invoke(state: GameState): Boolean
    // TODO add extra operator to discriminate for clusters. E.g: passing a ClusterState be able to guess if the filter will apply

    companion object {
        val ANY: GameEventFilter = object : GameEventFilter {
            override fun invoke(state: GameState): Boolean = true
            override fun toString(): String = "ANY"
        }
        val NONE: GameEventFilter = object : GameEventFilter {
            override fun invoke(state: GameState): Boolean = false
            override fun toString(): String = "NONE"}
        }

    data class Or(val filters: List<GameEventFilter>) : GameEventFilter {
        override fun invoke(state: GameState): Boolean = filters.any { it(state) }
        override fun toString(): String {
            val sb = StringBuilder()
            sb.append("OR(")
            sb.append(filters.joinToString(", ") { it.toString() })
            sb.append(")")
            return sb.toString()
        }
    }

    data class And(val filters: List<GameEventFilter>) : GameEventFilter {
        override fun invoke(state: GameState): Boolean = filters.all { it(state) }
        override fun toString(): String {
            val sb = StringBuilder()
            sb.append("AND(")
            sb.append(filters.joinToString(", ") { it.toString() })
            sb.append(")")
            return sb.toString()
        }
    }

    data class Not(val filter: GameEventFilter) : GameEventFilter {
        override fun invoke(state: GameState): Boolean = !filter(state)
        override fun toString(): String {
            val sb = StringBuilder()
            sb.append("NOT(")
            sb.append(filter)
            sb.append(")")
            return sb.toString()
        }
    }
}

