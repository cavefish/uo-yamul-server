package dev.cavefish.yamul

import org.springframework.boot.SpringBootConfiguration
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.ComponentScan
import org.springframework.test.context.ActiveProfiles

@ComponentScan(basePackages = ["dev.cavefish.yamul.backend"])
@SpringBootTest
@SpringBootConfiguration
@ActiveProfiles("test")
class IntegrationTest